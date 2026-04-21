# Day 11 - Worker Pool Pattern

## What I Learned

- Worker pool architecture
- Job queue using channels
- Controlled concurrency
- 👉 Why use it?
  Prevent resource exhaustion
  Control concurrency
  Efficient job processing
- 👉 Where used?
  background jobs
  message queues
  batch processing

## 📝 Notes

- Worker pools help process jobs efficiently without overwhelming the system.
- Architecture:
  Jobs Channel ---> Workers ---> Results Channel
- 👉 Think:
  Jobs = Kafka topic / Redis queue
  Workers = consumers
- ✅ Controlled concurrency
  Only 3 workers running at a time
- ✅ Queue-based processing
  Jobs pushed → workers consume
- ✅ Scalable pattern
  Increase workers → more throughput
- jobs channel = Kafka topic / Redis queue
- worker = consumer service
- results = processed output / DB write
- 👉 This is literally: Your BullMQ worker system in Go

## Practice

- Job processing system with 3 workers
