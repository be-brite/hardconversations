# AI Recruiter

Given a job description and a database of resumes, find the candidates that are most suited to the job and generate an email address to send to them.

See generated client usage in [main.go](https://github.com/be-brite/hardconversations/blob/main/samples/recruiter/main.go).

# hardc.yaml

```yaml
version: 1
conversations:
  - path: "./autorecruiter"
    instruction: |
      Given a list of resumes, you are able to determine which ones are the best fit for the job description.
    questions:
      - function_name: RankResumes
        prompt: Return just the IDs of between 1 and 3 resumes in a comma-separated list, ranked from best to worst fit for the job description. Do not include resumes that are not a good fit.
        input: "string"
        output: "[]int"

      - function_name: GetCandidateInfo
        prompt: "Return the candidate info from the resume"
        input: string
        output: github.com/be-brite/hardconversations/samples/recruiter/resumes.Candidate

      - function_name: GenerateRecruiterMessage
        prompt: Generate a message to send to the candidate about the job; mention what you like about their resume and why you think they would be a good fit for the job.
        input: "github.com/be-brite/hardconversations/samples/recruiter/resumes.RecruiterMessageRequest"
        output: github.com/be-brite/hardconversations/samples/recruiter/resumes.Email
```

## Usage

```go
rankedResumeIDs, _, err := t.RankResumes(ctx, developerJobDescription)
if err != nil {
  return
}

for _, resumeID := range rankedResumeIDs {
  ...
}
```
