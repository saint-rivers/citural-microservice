package com.saintrivers.commandtower.module.project.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import com.saintrivers.commandtower.shared.entity.Project;

public interface ProjectRepository extends JpaRepository<Project, Long> {

}
