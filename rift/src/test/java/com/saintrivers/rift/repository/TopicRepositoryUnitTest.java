package com.saintrivers.rift.repository;

import java.util.ArrayList;
import java.util.List;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.data.mongo.DataMongoTest;

import com.saintrivers.rift.domain.Course;
import com.saintrivers.rift.domain.Resource;
import com.saintrivers.rift.domain.Section;
import com.saintrivers.rift.domain.Topic;

@DataMongoTest
public class TopicRepositoryUnitTest {

    @Autowired
    private TopicRepository topicRepository;

    @Autowired
    private CourseRepository courseRepository;

    @Test
    void shouldCreateTopic() {
        // Topic topic = new Topic();
        // List<Section> sections = new ArrayList<>() {
        //     {
        //         Section section = new Section();
        //         section.setName("message brokers");
        //         section.setResearchPoints(List.of("pub/sub pattern", "message queues"));
        //         section.setResources(
        //                 List.of(new Resource("https://www.mongodb.com/compatibility/spring-boot", List.of("mongo"))));
        //         add(section);
        //     }
        // };

        // Course course = new Course();
        // course.setName("microservices");
        // topic.setCourseInformation(course);
        // topic.setSections(sections);

        // courseRepository.save(course);
        // topicRepository.save(topic);
    }
}
