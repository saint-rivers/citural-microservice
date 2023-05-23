package com.saintrivers.rift.domain;

import java.util.List;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class Section {
    private String name;
    private List<ResearchPoint> researchPoints;
    private List<Resource> resources;
}