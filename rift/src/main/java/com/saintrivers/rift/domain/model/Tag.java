package com.saintrivers.rift.domain.model;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import lombok.Data;

@Data
@Document("tag")
public class Tag {
    @Id
    private String id;
    private String name;
    private String type;
}
