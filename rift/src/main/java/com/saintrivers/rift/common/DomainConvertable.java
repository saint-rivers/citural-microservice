package com.saintrivers.rift.common;

public interface DomainConvertable<D, R> {
    D toDomain(R request);
}
