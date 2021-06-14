package bridge

import "fmt"

func Example_Bridge() {
	//Output:
	//hello tom,nice to meet you via email,it is common level
	//hello tom,it's on fire via sms,it is urgency level
	s1 := &SendCommonMessage{}
	res1 := s1.Send("tom", "nice to meet you", &MessageEmail{})

	s2 := &SendUrgencyMessage{}
	res2 := s2.Send("tom", "it's on fire", &MessageSMS{})

	fmt.Println(res1)
	fmt.Println(res2)
}
