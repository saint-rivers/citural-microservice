package com.saintrivers.commandtower.module.project.service;

import java.util.List;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Sort;
import org.springframework.stereotype.Service;

import com.saintrivers.commandtower.module.project.mapper.EndpointMapper;
import com.saintrivers.commandtower.module.project.mapper.ProjectMapper;
import com.saintrivers.commandtower.module.project.repository.ProjectRepository;
import com.saintrivers.commandtower.module.project.web.request.EndpointRequest;
import com.saintrivers.commandtower.module.project.web.request.ProjectRequest;
import com.saintrivers.commandtower.module.project.web.response.EndpointResponse;
import com.saintrivers.commandtower.module.project.web.response.ProjectResponse;
import com.saintrivers.commandtower.shared.entity.Endpoint;
import com.saintrivers.commandtower.shared.entity.Project;
import com.saintrivers.commandtower.shared.response.PagedResponse;

import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor
public class ProjectServiceImpl implements ProjectService {

    private final ProjectRepository projectRepository;
    private final ProjectMapper projectMapper;
    private final EndpointMapper endpointMapper;

    @Override
    public void create(ProjectRequest r) {
        var domain = projectMapper.toEntity(r);
        projectRepository.save(domain);
    }

    @Override
    public PagedResponse<List<ProjectResponse>> getPagedData(int page, int size) {
        PageRequest pageable = PageRequest.of(page, size, Sort.Direction.DESC, "createdAt");
        Page<Project> projects = projectRepository.findAll(pageable);
        Page<ProjectResponse> data = projects.map((p) -> {
            return projectMapper.toResponse(p);
        });
        return new PagedResponse<>(data.getContent(), data.getTotalPages());
    }

    @Override
    public void addEndpoint(EndpointRequest endpointRequest) {
        Project existingProject = projectRepository
                .findById(endpointRequest.projectId())
                .orElseThrow();
        Endpoint endpoint = endpointMapper.toEntity(endpointRequest);
        existingProject.getEndpoints().add(endpoint);
        projectRepository.save(existingProject);
    }

    @Override
    public List<EndpointResponse> getEndpoints(long projectId) {
        var selectedProject = projectRepository.findById(projectId).orElseThrow();
        List<Endpoint> endpoints = selectedProject.getEndpoints();
        return endpointMapper.toResponse(endpoints);
    }
}
