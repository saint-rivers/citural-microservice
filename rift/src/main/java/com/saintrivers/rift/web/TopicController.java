package com.saintrivers.rift.web;

import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.saintrivers.rift.app.topic.TopicService;
import com.saintrivers.rift.common.response.BaseResponse;
import com.saintrivers.rift.common.response.BaseResponseBuilder;
import com.saintrivers.rift.web.request.TopicRequest;

import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.PathVariable;

@RestController
@RequestMapping("/api/v1/topics")
@RequiredArgsConstructor
public class TopicController {

    private final TopicService topicService;

    @PostMapping
    public BaseResponse<?> test(@RequestBody TopicRequest topic) {
        topicService.create(topic);
        return BaseResponseBuilder.of("topic created");
    }

    @GetMapping
    public BaseResponse<?> getTopicsOfCourse(@RequestParam String courseId) {
        var payload = topicService.getByCourseId(courseId);
        return BaseResponseBuilder.of("data fetched", payload);
    }

    @PutMapping("/{topicId}")
    public BaseResponse<?> updateTopic(@PathVariable String topicId, @RequestBody TopicRequest topic) {
        topicService.update(topicId, topic);
        return BaseResponseBuilder.of("topic updated");
    }

    @DeleteMapping("/{topicId}")
    public BaseResponse<?> deleteTopic(@PathVariable String topicId) {
        topicService.delete(topicId);
        return BaseResponseBuilder.of("topic updated");
    }
}
