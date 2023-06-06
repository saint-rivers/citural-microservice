package com.saintrivers.commandtower.module.project.repository;

import java.util.UUID;

import org.springframework.data.jpa.repository.JpaRepository;

import com.saintrivers.commandtower.shared.entity.Endpoint;

public interface EndpointRepository extends JpaRepository<Endpoint, UUID> {
    
}
