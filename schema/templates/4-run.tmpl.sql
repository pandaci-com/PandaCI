CREATE TYPE run_status_enum AS ENUM ('queued', 'running', 'completed', 'pending', 'awaiting-approval');
CREATE TYPE run_conclusion_enum AS ENUM ('success', 'failure', 'cancelled', 'skipped');

CREATE TYPE run_trigger_enum AS ENUM ('push', 'pull_request-opened', 'pull_request-synchronize', 'pull_request-closed', 'manual');

CREATE TABLE workflow_run (
    id CHAR(21) NOT NULL PRIMARY KEY,
    number INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    finished_at TIMESTAMP,
    project_id CHAR(21) NOT NULL,
    status run_status_enum NOT NULL,
    conclusion run_conclusion_enum,
    name VARCHAR(255) NOT NULL,
    git_title VARCHAR(255),
    runner VARCHAR(255) NOT NULL DEFAULT 'ubuntu-2x',
    git_sha TEXT NOT NULL,
    git_branch TEXT NOT NULL,
    committer_email TEXT,
    user_id UUID, -- This gets populated when we know the user
    pr_number INT,
    build_minutes INT NOT NULL DEFAULT 0,
    trigger run_trigger_enum NOT NULL DEFAULT 'push',
    alerts JSONB DEFAULT '[]'::JSONB, -- { type: 'error' | 'warning', title: string, message: string }[]
    FOREIGN KEY (project_id) REFERENCES project (id) ON DELETE CASCADE,
    UNIQUE (project_id, number)
);

CREATE TABLE job_run (
    id CHAR(21) NOT NULL PRIMARY KEY,
    number INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    finished_at TIMESTAMP,
    workflow_run_id CHAR(21) NOT NULL,
    status run_status_enum NOT NULL,
    build_minutes INT NOT NULL DEFAULT 0,
    conclusion run_conclusion_enum,
    name VARCHAR(255) NOT NULL,
    runner VARCHAR(255) NOT NULL DEFAULT 'ubuntu-4x',
    FOREIGN KEY (workflow_run_id) REFERENCES workflow_run (id) ON DELETE CASCADE,
    UNIQUE (workflow_run_id, number)
);

CREATE TABLE task_run (
    id CHAR(21) NOT NULL PRIMARY KEY,
    workflow_run_id CHAR(21) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    finished_at TIMESTAMP,
    job_run_id CHAR(21) NOT NULL,
    name VARCHAR(255) NOT NULL,
    status run_status_enum NOT NULL,
    conclusion run_conclusion_enum,
    docker_image VARCHAR(255),
    meta JSONB, -- key-value pairs, these are then displayed in the UI
    FOREIGN KEY (job_run_id) REFERENCES job_run (id) ON DELETE CASCADE,
    FOREIGN KEY (workflow_run_id) REFERENCES workflow_run (id) ON DELETE CASCADE
);

CREATE TYPE step_type_enum AS ENUM ('exec');

CREATE TABLE step_run (
    id CHAR(21) NOT NULL PRIMARY KEY,
    workflow_run_id CHAR(21) NOT NULL,
    type step_type_enum NOT NULL,
    created_at TIMESTAMP NOT NULL,
    finished_at TIMESTAMP,
    task_run_id CHAR(21) NOT NULL,
    job_run_id CHAR(21) NOT NULL,
    status run_status_enum NOT NULL,
    conclusion run_conclusion_enum,
    name VARCHAR(255) NOT NULL, -- For exec step, we use the command as the name, if its too long, we truncate it
    meta JSONB, -- key-value pairs, these are then displayed in the UI
    FOREIGN KEY (task_run_id) REFERENCES task_run (id) ON DELETE CASCADE,
    FOREIGN KEY (job_run_id) REFERENCES job_run (id) ON DELETE CASCADE,
    FOREIGN KEY (workflow_run_id) REFERENCES workflow_run (id) ON DELETE CASCADE
);
