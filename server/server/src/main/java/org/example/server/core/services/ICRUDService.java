package org.example.server.core.services;

import java.util.Collection;
import java.util.NoSuchElementException;

public interface ICRUDService<T, ID> {

    T findById(ID id) throws NoSuchElementException;

    Collection<T> findAll();

    Collection<T> findAllByIds(Iterable<ID> ids);

    T save(T entity);

    void deleteById(ID id);

    void delete(T entity);

}