package main

func limit(result string) string {
	if len(result) > 40 {
		return result[:40] + "..."
	}
	return result
}
