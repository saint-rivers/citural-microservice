package com.saintrivers.commandtower.module.project.web.controller;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.saintrivers.commandtower.module.project.service.ProjectService;
import com.saintrivers.commandtower.module.project.web.request.EndpointRequest;

import lombok.RequiredArgsConstructor;

@RestController
@RequiredArgsConstructor
@RequestMapping("/api/v1/endpoints")
@CrossOrigin
public class EndpointController {

    private final ProjectService projectService;

    @PostMapping
    public ResponseEntity<?> createEndpoint(@RequestBody EndpointRequest endpointRequest) {
        projectService.addEndpoint(endpointRequest);
        return ResponseEntity.ok().build();
    }
}
