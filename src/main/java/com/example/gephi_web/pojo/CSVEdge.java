package com.example.gephi_web.pojo;

import lombok.*;

@Data
@Getter
@Setter
@ToString
@EqualsAndHashCode
public class CSVEdge {
    Integer id;
    String source;
    String target;
    String attributes;

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public String getSource() {
        return source;
    }

    public void setSource(String source) {
        this.source = source;
    }

    public String getTarget() {
        return target;
    }

    public void setTarget(String target) {
        this.target = target;
    }

    public String getAttributes() {
        return attributes;
    }

    public void setAttributes(String attributes) {
        this.attributes = attributes;
    }
}
