package com.saintrivers.rift.controller;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.saintrivers.rift.domain.Topic;
import com.saintrivers.rift.repository.TopicRepository;

import io.swagger.v3.oas.annotations.parameters.RequestBody;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;

@RestController
@RequestMapping("/api/v1/topics")
@RequiredArgsConstructor
public class TopicController {
    
    private final TopicRepository topicRepository;
    
    @PostMapping
    public void test(@Valid @RequestBody Topic topic){
        topicRepository.save(topic);
    }
}
