import {Attr, Attribute, BelongsTo, Boolean, HasMany, HasOne, Model, MorphMany, MorphOne, MorphTo, Number} from '@vuex-orm/core';
import {omit, pullAll} from 'lodash-es';

export class BaseModel extends Model {
  static primaryKey = '_id';

  static state() {
    return {};
  }

  static apiConfig = {
    baseURL: 'http://localhost:9000/api',
  }

  static fieldsKeys() {
    return Object.keys(omit(this.fields(), this.primaryKey));
  }

  static relationFields() {
    return this.fieldsKeys().reduce((list, field) => {
      let fieldAttribute = this.fields()[field];
      if (this.isFieldRelation(fieldAttribute)) {
        list.push(`${field}_id`);
        list.push(field);
      }
      return list;
    }, []);
  }

  static pagination(page, size) {
    const total = this.query().count();
    const pageNum = (page && parseInt(page)) || 1;
    const pageSize = (size && parseInt(size)) || 10;
    const totalPage = Math.ceil(total / pageSize) || 0;
    const offset = (pageNum - 1) * pageSize || 0;
    const next = (pageNum >= totalPage ? total % pageSize : pageSize) + 1;
    return {
      total,
      pageNum,
      pageSize,
      totalPage,
      offset,
      next,
    };
  }

  static pageQuery(pagination, query) {
    if (!query) query = this.query();
    return query.offset(pagination.offset).limit(pagination.pageSize);
  }

  static nonRelationFields() {
    return pullAll(this.fieldsKeys(), this.relationFields());
  }

  static isFieldAttribute(field) {
    return (
        field instanceof Attr ||
        field instanceof String ||
        field instanceof Number ||
        field instanceof Boolean
    );
  }

  static isFieldRelation(field) {
    return (
        field instanceof BelongsTo ||
        field instanceof HasOne ||
        field instanceof HasMany ||
        field instanceof MorphTo ||
        field instanceof MorphOne ||
        field instanceof MorphMany
    );
  }

  skipField(field) {
    let shouldSkipField = false;
    this.getRelations().forEach(relation => {
      if (
          (relation instanceof BelongsTo || relation instanceof HasOne) &&
          relation.foreignKey === field
      ) {
        shouldSkipField = true;
        return false;
      }
      return true;
    });

    return shouldSkipField;
  }

  /**
   * Determines if we should eager load (means: add as a field in the graphql query) a related entity. belongsTo or
   * hasOne related entities are always eager loaded. Others can be added to the `eagerLoad` array of the model.
   *
   * @param {string} fieldName Name of the field
   * @param {Attribute} field Relation field
   * @param {Model} relatedModel Related model
   * @returns {boolean}
   */
  shouldEagerLoadRelation(fieldName, field, relatedModel) {
    if (field instanceof HasOne || field instanceof BelongsTo || field instanceof MorphOne) {
      return true;
    }

    const eagerLoadList = this.eagerLoad || [];
    return (
        eagerLoadList.find(n => {
          return n === relatedModel.singularName || n === relatedModel.pluralName || n === fieldName;
        }) !== undefined
    );
  }

  getRelations() {
    return this.fields().reduce(function(relations, field, name) {
      if (!this.isFieldAttribute(field)) {
        relations.set(name, field);
      }
      return relations;
    }, []);
  }
}
