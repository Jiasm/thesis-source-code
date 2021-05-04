'use strict'

import * as util from '../src/util/index.js'
import { expect } from 'chai'

describe('frontend uint test', function() {
  describe('check status map entity validate', function() {
    Object.values(util.statusMap).forEach((value, index) => {
      it(`${index} equals "${value}"`, function() {
        expect(util.getStatus(index)).to.be.equal(value)
      })
    })
    it(`default equals "${util.statusMap[0]}"`, function() {
      expect(util.getStatus(999)).to.be.equal(util.statusMap[0])
    })
  })
  describe('check priority map entity validate', function() {
    Object.values(util.priorityMap).forEach((value, index) => {
      it(`${index} equals "${value}"`, function() {
        expect(util.getPriority(index)).to.be.equal(value)
      })
    })
    it(`default equals "${util.priorityMap[0]}"`, function() {
      expect(util.getPriority(999)).to.be.equal(util.priorityMap[0])
    })
  })
  describe('check type map entity validate', function() {
    Object.values(util.typeMap).forEach((value, index) => {
      it(`${index} equals "${value}"`, function() {
        expect(util.getTaskType(index)).to.be.equal(value)
      })
    })
    it(`default equals "${util.typeMap[0]}"`, function() {
      expect(util.getTaskType(999)).to.be.equal(util.typeMap[0])
    })
  })
  describe('check role map entity validate', function() {
    Object.values(util.roleMap).forEach((value, index) => {
      it(`${index} equals "${value}"`, function() {
        expect(util.getRole(index)).to.be.equal(value)
      })
    })
    it(`default equals "${util.roleMap[0]}"`, function() {
      expect(util.getRole(999)).to.be.equal(util.roleMap[0])
    })
  })
  describe('check userStatus map entity validate', function() {
    Object.values(util.userStatusMap).forEach((value, index) => {
      it(`${index} equals "${value}"`, function() {
        expect(util.getUserStatus(index)).to.be.equal(value)
      })
    })
    it(`default equals "${util.userStatusMap[0]}"`, function() {
      expect(util.getUserStatus(999)).to.be.equal(util.userStatusMap[0])
    })
  })
  describe('build selector', function() {
    it(`{ a: 1, b: 2 } will transform to [{ value: 1, label: "a" }, { value: 2, label: "b" }]`, function() {
      expect(JSON.stringify(util.buildSelector({ 1: "a", 2: "b" }))).to.be.equal(JSON.stringify([{ value: 1, label: "a" }, { value: 2, label: "b" }]))
    })
  })
  describe('date formate', function() {
    it(`1620128405 will transform to "2021-05-04 19:40:00"`, function() {
      expect(util.formatDate(1620128405)).to.be.equal('2021-05-04 19:40:00')
    })
  })
})