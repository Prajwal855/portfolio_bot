package main

import "strings"

var commandResponses = map[string]string{
	"/start": strings.TrimSpace(`Hi, I'm Prajwal's portfolio bot.

Use these commands:
/help - List all commands
/about - Professional summary
/skills - Core skills
/experience - Work experience
/projects - Key projects
/contact - Contact links`),

	"/help": strings.TrimSpace(`Available commands:
/start
/about
/skills
/experience
/projects
/contact
/website
/github
/linkedin`),

	"/about": strings.TrimSpace(`Backend Engineer with 4+ years of experience building scalable, high-reliability systems using Ruby on Rails and PostgreSQL.

Currently working onsite at Razorpay (Engage BU), delivering payments, wallets, loyalty, and rewards platforms for enterprise clients including Audi, Visa, HDFC, Yes Bank, and LTFS.`),

	"/skills": strings.TrimSpace(`Languages: Ruby, Python, JavaScript, TypeScript
Backend: Ruby on Rails, REST APIs, SSO, Wallet Systems, Sidekiq
Databases: PostgreSQL, SQL
Infra/Tools: Git, Docker, Kubernetes, Linux, Jenkins
Concepts: Distributed Systems, API Design, Integrations, Config-driven Platforms, Reliability`),

	"/experience": strings.TrimSpace(`Senior Engineer (Backend - Ruby on Rails)
Happiest Minds Technologies (Onsite at Razorpay, Engage BU)
2024 - Present | Bengaluru

Software Engineer
CognitiveClouds Software Pvt Ltd
Feb 2022 - 2024 | Bengaluru`),

	"/projects": strings.TrimSpace(`Key projects:
1) Audi Loyalty Program - SSO, wallets, transactions, OTP, payments
2) PCI-DSS + Rails Upgrade - 300+ vulns remediated across 9+ services
3) HDFC SI Fulfillment - Sidekiq retries + parent-child order model
4) Yes Bank Rewards - scheduled notification worker
5) LTFS Consumer Loyalty - SSO + wallet integration + tech spec ownership
6) Visa projects (Hajj/IPL/Cambodia) - platformization and rapid delivery`),

	"/contact": strings.TrimSpace(`Phone: +91 9900717474
Email: prajwal.mysore0077@gmail.com
Website: https://prajwal-portfolio.com
GitHub: https://github.com/Prajwal855
LinkedIn: https://linkedin.com/in/prajwal-m`),

	"/website":  "https://prajwal-portfolio.com",
	"/github":   "https://github.com/Prajwal855",
	"/linkedin": "https://linkedin.com/in/prajwal-m",
}

func responseFor(text string) string {
	text = strings.TrimSpace(strings.ToLower(text))

	if resp, ok := commandResponses[text]; ok {
		return resp
	}

	return "Unknown command. Use /help to see available commands."
}
