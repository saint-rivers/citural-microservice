package com.saintrivers.commandtower.module.project.web.controller;

import java.util.List;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.saintrivers.commandtower.module.project.service.ProjectService;
import com.saintrivers.commandtower.module.project.web.request.ProjectRequest;
import com.saintrivers.commandtower.module.project.web.response.EndpointResponse;
import com.saintrivers.commandtower.shared.response.PagedResponse;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;

@RestController
@RequiredArgsConstructor
@RequestMapping("/api/v1/projects")
@CrossOrigin
@Slf4j
public class ProjectController {

    private final ProjectService projectService;

    @PostMapping
    public ResponseEntity<?> create(@RequestBody ProjectRequest projectRequest) {
        projectService.create(projectRequest);
        return ResponseEntity.ok().build();
    }

    @GetMapping
    public PagedResponse<?> get(
            @RequestParam(defaultValue = "1") int page,
            @RequestParam(defaultValue = "5") int size) {
        return projectService.getPagedData(page - 1, size);
    }

    @GetMapping("/{projectId}/endpoints")
    public List<EndpointResponse> getEndpoints(@PathVariable int projectId) {
        return projectService.getEndpoints(projectId);
    }

}
