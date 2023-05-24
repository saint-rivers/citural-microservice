package com.saintrivers.rift.app.topic;

import java.util.List;

import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.data.mongodb.core.query.Update;
import org.springframework.stereotype.Service;

import com.saintrivers.rift.domain.Topic;
import com.saintrivers.rift.web.request.TopicRequest;

import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor
public class TopicServiceImpl implements TopicService {

    private final TopicRepository topicRepository;
    private final TopicMapper topicMapper;
    private final MongoTemplate mongoTemplate;

    @Override
    public void create(TopicRequest topicRequest) {
        Topic topic = topicMapper.toDomain(topicRequest);
        topic.setIsDeleted(false);
        topicRepository.save(topic);
    }

    @Override
    public List<Topic> getByCourseId(String courseId) {
        return topicRepository.selectWhereCourseIdIs(courseId);
    }

    @Override
    public void delete(String topicId) {
        Query query = new Query();
        query.addCriteria(Criteria.where("_id").is(topicId));
        Topic topic = mongoTemplate.findOne(query, Topic.class);

        topic = topicMapper.prepareUpdate(topic, new Topic());
        topic.setIsDeleted(true);

        applyUpdates(query, topicId, topic);
    }

    @Override
    public void update(String topicId, TopicRequest topicRequest) {
        Query query = new Query();
        query.addCriteria(Criteria.where("_id").is(topicId));
        Topic topic = mongoTemplate.findOne(query, Topic.class);

        Topic updateRequest = topicMapper.toDomain(topicRequest);
        topic = topicMapper.prepareUpdate(topic, updateRequest);

        applyUpdates(query, topicId, topic);
    }

    private void applyUpdates(Query query, String topicId, Topic topic) {
        Update update = new Update();
        update.set("changes", topic.getChanges());
        update.set("sections", topic.getSections());
        update.set("lastUpdated", topic.getLastUpdated());
        update.set("index", topic.getIndex());
        update.set("isDeleted", topic.getIsDeleted());
        mongoTemplate.findAndModify(query, update, Topic.class);
    }

}
