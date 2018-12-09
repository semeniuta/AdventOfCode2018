package main

import (
	"fmt"
	"github.com/semeniuta/AdventOfCode2018/aoccommons"
)

type task string

type taskInfo struct {
	dependsOn []task
	next []task
}

type schedule struct {
	tasks map[task]*taskInfo
	available map[task]bool
	done map[task]bool
}

func makeSchedule() schedule {

	var s schedule
	s.tasks = make(map[task]*taskInfo)
	s.available = make(map[task]bool)
	s.done = make(map[task]bool)

	return s
}

func (s *schedule) taskHasNoDependencies(t task) bool {
	return len(s.tasks[t].dependsOn) == 0
}

func (s *schedule) addTask(t task) {
	
	_, ok := s.tasks[t]
	if !ok {
		s.tasks[t] = &taskInfo{}
	}
	
}

func (s *schedule) addDependency(t task, tPrev task) {
	
	s.tasks[t].dependsOn = append(s.tasks[t].dependsOn, tPrev)
	s.tasks[tPrev].next = append(s.tasks[tPrev].next, t)
}

func (s *schedule) initialize() {
	for t := range s.tasks {
		if s.taskHasNoDependencies(t) {
			s.available[t] = true
		}
	}
}

func (s *schedule) nextTask() task {
	
	if len(s.available) == 0 {
		return ""
	}

	smallest := task("z")
	for t := range s.available {
		if t < smallest {
			smallest = t
		}
	}

	return smallest

}

func (s *schedule) executeNext() task {

	t := s.nextTask()
	if t == "" {
		return t
	}

	s.done[t] = true

	info := s.tasks[t]
	delete(s.available, t)

	for _, tNext := range info.next {
		
		allDepsDone := true
		
		for _, dep := range s.tasks[tNext].dependsOn {
			_, ok := s.done[dep]
			if !ok {
				allDepsDone = false
				break
			}
		}

		if allDepsDone {
			s.available[tNext] = true
		}
		
	}

	return t
}

func (s *schedule) run() string {

	res := ""

	for true {

		//fmt.Println("Available:", s.available, ". Done:", s.done)

		tx := s.executeNext()

		if tx == "" {
			break
		}

		res += string(tx)
	} 

	return res
}

func parseInput(filename string) schedule {

	s := makeSchedule()

	scanner := aoccommons.CreateScanner(filename)

	re := aoccommons.CompileRegex("Step (?P<before>[A-Z]) must be finished before step (?P<after>[A-Z]) can begin.")

	for scanner.Scan() {
		line := scanner.Text()

		d := aoccommons.RegexParse(re, line)

		tBefore := task(d["before"])
		tAfter := task(d["after"])

		s.addTask(tBefore)
		s.addTask(tAfter)
		s.addDependency(tAfter, tBefore)
	}

	s.initialize()

	return s

}

func main() {

	// Correct (1st): GNJOCHKSWTFMXLYDZABIREPVUQ

	s := parseInput("input.txt")
	seq := s.run()
	fmt.Println("Sequence:", seq)
}