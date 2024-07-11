package org.example.server.core.services.implementation;

import org.example.server.core.services.ICRUDService;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Collection;
import java.util.NoSuchElementException;

public abstract class CRUDService<T, ID> implements ICRUDService<T, ID> {

    private JpaRepository<T, ID> repository;

    public CRUDService(JpaRepository<T, ID> repository) {
        super();
        this.repository = repository;
    }

    public T findById(ID id) throws NoSuchElementException {
        return repository.findById(id).orElseThrow();
    }

    public Collection<T> findAll() {
        return repository.findAll();
    }

    public Collection<T> findAllByIds(Iterable<ID> ids) {
        return repository.findAllById(ids);
    }

    public T save(T entity) {
        return repository.save(entity);
    }

    public void deleteById(ID id) {
        repository.deleteById(id);
    }

    public void delete(T entity) {
        repository.delete(entity);
    }

}
