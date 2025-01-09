package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/SHIVAM-GOUR/social_go_app/internal/store"
)

var usernames = []string{
	"Alice", "Bob", "Charlie", "David", "Eve",
	"Frank", "Grace", "Hank", "Ivy", "Jack",
	"Karen", "Liam", "Mona", "Nate", "Olivia",
	"Paul", "Quinn", "Rachel", "Steve", "Tina",
	"Uma", "Victor", "Wendy", "Xander", "Yara",
	"Zane", "Amelia", "Benjamin", "Chloe", "Daniel",
	"Eliza", "Finn", "Gabriella", "Henry", "Isla",
	"James", "Kylie", "Lucas", "Mia", "Nathan",
	"Oscar", "Piper", "Quincy", "Ruby", "Samuel",
	"Taylor", "Uma", "Vincent", "Willow", "Xavier",
}

var titles = []string{
	"10 Tips for Productivity",
	"Mastering Minimalism",
	"How to Stay Motivated",
	"Essential Coding Practices",
	"Healthy Eating on a Budget",
	"The Art of Mindfulness",
	"Travel Hacks for Beginners",
	"Building a Personal Brand",
	"Understanding Cryptocurrency",
	"Secrets to Better Sleep",
	"DIY Home Decor Ideas",
	"Fitness Myths Debunked",
	"Creating a Capsule Wardrobe",
	"Effective Time Management",
	"Starting Your First Blog",
	"Learn Photography Basics",
	"Improving Mental Health",
	"Top 5 Books to Read",
	"Cooking Made Simple",
	"Exploring Nature Trails",
}

var contents = []string{
	"Discover 10 actionable tips to boost your daily productivity.",
	"Learn how minimalism can simplify your life and mindset.",
	"Practical strategies to stay motivated in challenging times.",
	"A beginner's guide to writing clean and maintainable code.",
	"Explore healthy eating habits without breaking the bank.",
	"Learn how to embrace mindfulness for a stress-free life.",
	"Top travel hacks to save money and enhance your adventures.",
	"Steps to build a personal brand and grow your online presence.",
	"A simple guide to understanding how cryptocurrencies work.",
	"Proven methods to improve your sleep quality tonight.",
	"Easy and creative DIY decor ideas to transform your home.",
	"Debunking common fitness myths for a healthier lifestyle.",
	"Tips to create a capsule wardrobe and simplify your mornings.",
	"Time management techniques to get more done in less time.",
	"A beginner-friendly guide to launching your first blog.",
	"Master the basics of photography and capture stunning shots.",
	"Simple habits to improve your mental health every day.",
	"A curated list of 5 must-read books to inspire and educate.",
	"Quick and easy cooking tips for delicious homemade meals.",
	"Explore the beauty of nature trails for fitness and serenity.",
}

var tags = []string{
	"Lifestyle", "Productivity", "Technology", "Travel",
	"Health", "Fitness", "DIY", "Food",
	"Finance", "Parenting", "Mindfulness", "Photography",
	"Education", "Fashion", "Gaming", "Coding",
	"Books", "Nature", "Marketing", "Entrepreneurship",
}

var comments = []string{
	"Great insights, thanks for sharing!",
	"This is exactly what I needed today.",
	"Can you elaborate more on this topic?",
	"I totally agree with your point of view.",
	"Thanks for breaking this down so clearly!",
	"This was super helpful, keep it up!",
	"Interesting perspective, but I have a different take.",
	"Where can I find more details about this?",
	"Amazing post! Learned a lot.",
	"Can you recommend resources to dive deeper?",
	"Thanks for simplifying such a complex topic.",
	"This is very well-written and informative.",
	"Looking forward to reading more from you!",
	"Have you tried implementing this in real life?",
	"Your examples make this so much easier to understand.",
	"This has inspired me to take action!",
	"I have a question about one of your points.",
	"Excellent advice, I'll apply this immediately.",
	"Thanks for your hard work on this post!",
	"Do you have any tips for beginners like me?",
}

func Seed(store *store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user: ", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post: ", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating post: ", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123123",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		post := posts[rand.Intn(len(posts))]

		cms[i] = &store.Comment{
			PostID:  post.ID,
			UserID:  user.ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return cms
}
