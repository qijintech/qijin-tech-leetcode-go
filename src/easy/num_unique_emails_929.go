func numUniqueEmails(emails []string) int {
	emailSet := make(map[string]bool,len(emails))
	res := 0
	for _, email := range emails {
		emailSlice := strings.Split(email, "@")
		localName := strings.Replace(strings.Split(emailSlice[0], "+")[0], ".","",-1)
		formatedEmail := localName + "@" + emailSlice[1]
		if emailSet[formatedEmail] {
			continue
		}else {
			emailSet[formatedEmail] = true
			res ++
		}
	}
	return res
}