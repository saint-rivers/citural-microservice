package com.saintrivers.rift.domain;

import java.util.List;

import org.springframework.data.annotation.CreatedBy;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import com.saintrivers.rift.common.Domain;
import com.saintrivers.rift.domain.model.Section;

import lombok.Data;
import lombok.EqualsAndHashCode;

@Document(collection = "topic")
@Data
@EqualsAndHashCode(callSuper = false)
public class Topic extends Domain<Topic> {

    @Id
    private String id;
    private List<Section> sections;
    private Integer index = 0;
    private String courseId;

    @CreatedBy
    private String userId;

}
