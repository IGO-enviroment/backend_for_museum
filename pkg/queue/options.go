package queue

// Option -.
type Option func(*Queue)

// // MaxPoolSize -.
// func MaxPoolSize(size int) Option {
// 	return func(c *Queue) {
// 		c.maxPoolSize = size
// 	}
// }

// // ConnAttempts -.
// func ConnAttempts(attempts int) Option {
// 	return func(c *Queue) {
// 		c.connAttempts = attempts
// 	}
// }

// // ConnTimeout -.
// func ConnTimeout(timeout time.Duration) Option {
// 	return func(c *Queue) {
// 		c.connTimeout = timeout
// 	}
// }
