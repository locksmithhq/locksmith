ALTER TABLE project_acl DROP CONSTRAINT project_acl_role_id_fkey;
ALTER TABLE project_acl ADD CONSTRAINT project_acl_role_id_fkey
    FOREIGN KEY (role_id) REFERENCES roles(id);

ALTER TABLE project_acl DROP CONSTRAINT project_acl_module_id_fkey;
ALTER TABLE project_acl ADD CONSTRAINT project_acl_module_id_fkey
    FOREIGN KEY (module_id) REFERENCES modules(id);

ALTER TABLE project_acl DROP CONSTRAINT project_acl_action_id_fkey;
ALTER TABLE project_acl ADD CONSTRAINT project_acl_action_id_fkey
    FOREIGN KEY (action_id) REFERENCES actions(id);

ALTER TABLE project_acl DROP CONSTRAINT project_acl_project_id_fkey;
ALTER TABLE project_acl ADD CONSTRAINT project_acl_project_id_fkey
    FOREIGN KEY (project_id) REFERENCES projects(id);
