all: mcqTest
	8l -o goMcQueryTest goMcQueryTest.8

goMcQuery: goMcQuery.go commons
	8g goMcQuery.go
	
mcqTest: goMcQueryTest.go commons goMcQuery
	8g goMcQueryTest.go
	
commons: commons.go
	8g commons.go

clean: 
	rm *.8
