package main

import (
  "testing"
)

func TestMainArgumentsBad(t *testing.T) {
  payload := []string{"new"}
  status, msg := cli(payload)
  if status != 1 || msg != "Error: 3 arguments needed" {
    t.Errorf("Should return negative status with error instead return %d with message: %s", status, msg)
  }
}

func TestCommandBad(t *testing.T) {
  payload := []string{"new", "2", "warrior"}
  status, msg := cli(payload)
  if status != 1 || msg != "Error: Use create statement" {
    t.Errorf("Should return negative status with error instead return %d and message: %s", status, msg)
  }
}

func TestNumberOfInstancesBad(t *testing.T) {
  payload := []string{"create", "-2", "warrior"}
  status, msg := cli(payload)
  if status != 1 || msg != "Error: Invalid number of instances" {
    t.Errorf("Should return negative status with error instead return %d with message: %s", status, msg)
  }
}

func TestInstancePlanBad(t *testing.T) {
  payload := []string{"create", "1", "mortar"}
  status, msg := cli(payload)
  if status != 1 || msg != "Error: This plan is not supported" {
    t.Errorf("Should return negative status with error instead return %d with message: %s", status, msg)
  }
}

func TestInstancePlanWarriorOk(t *testing.T) {
  payload := []string{"create", "1", "warrior"}
  status, msg := cli(payload)
  if status != 0 || msg != "Success: Warrior created" {
    t.Errorf("Should return positive status without error instead return %d with message: %s", status, msg)
  }
}

func TestInstancePlanArcherOk(t *testing.T) {
  payload := []string{"create", "1", "archer"}
  status, msg := cli(payload)
  if status != 0 || msg != "Success: Archer created" {
    t.Errorf("Should return positive status without error instead return %d with message: %s", status, msg)
  }
}



