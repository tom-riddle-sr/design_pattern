package main

import "fmt"

// ============ Observer Interface ============

// Observer is the interface that all observers must implement
type Observer interface {
	Update(channelName string, videoTitle string)
	GetName() string
}

// ============ Subject Interface ============

// Subject is the interface for the publisher
type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify(videoTitle string)
}

// ============ Concrete Subject ============

// YouTubeChannel is the concrete subject (publisher)
type YouTubeChannel struct {
	channelName string
	subscribers []Observer
}

func NewYouTubeChannel(name string) *YouTubeChannel {
	return &YouTubeChannel{
		channelName: name,
		subscribers: make([]Observer, 0),
	}
}

func (y *YouTubeChannel) Attach(observer Observer) {
	y.subscribers = append(y.subscribers, observer)
	fmt.Printf("🔔 %s subscribed to %s\n", observer.GetName(), y.channelName)
}

func (y *YouTubeChannel) Detach(observer Observer) {
	for i, sub := range y.subscribers {
		if sub.GetName() == observer.GetName() {
			y.subscribers = append(y.subscribers[:i], y.subscribers[i+1:]...)
			fmt.Printf("🔕 %s unsubscribed from %s\n", observer.GetName(), y.channelName)
			return
		}
	}
}

func (y *YouTubeChannel) Notify(videoTitle string) {
	fmt.Printf("\n📢 [%s] uploaded a new video: \"%s\"\n", y.channelName, videoTitle)
	fmt.Println("Notifying all subscribers...")
	for _, subscriber := range y.subscribers {
		subscriber.Update(y.channelName, videoTitle)
	}
}

func (y *YouTubeChannel) UploadVideo(title string) {
	fmt.Printf("\n🎬 [%s] is uploading: \"%s\"\n", y.channelName, title)
	y.Notify(title)
}

// ============ Concrete Observers ============

// Subscriber is a concrete observer
type Subscriber struct {
	name string
}

func NewSubscriber(name string) *Subscriber {
	return &Subscriber{name: name}
}

func (s *Subscriber) GetName() string {
	return s.name
}

func (s *Subscriber) Update(channelName string, videoTitle string) {
	fmt.Printf("   📨 %s received notification: New video from [%s] - \"%s\"\n", s.name, channelName, videoTitle)
}

// EmailSubscriber - receives email notifications
type EmailSubscriber struct {
	name  string
	email string
}

func NewEmailSubscriber(name string, email string) *EmailSubscriber {
	return &EmailSubscriber{
		name:  name,
		email: email,
	}
}

func (e *EmailSubscriber) GetName() string {
	return e.name
}

func (e *EmailSubscriber) Update(channelName string, videoTitle string) {
	fmt.Printf("   📧 Email sent to %s (%s): New video from [%s] - \"%s\"\n", e.name, e.email, channelName, videoTitle)
}

// MobileSubscriber - receives mobile push notifications
type MobileSubscriber struct {
	name   string
	device string
}

func NewMobileSubscriber(name string, device string) *MobileSubscriber {
	return &MobileSubscriber{
		name:   name,
		device: device,
	}
}

func (m *MobileSubscriber) GetName() string {
	return m.name
}

func (m *MobileSubscriber) Update(channelName string, videoTitle string) {
	fmt.Printf("   📱 Push notification to %s (%s): New video from [%s] - \"%s\"\n", m.name, m.device, channelName, videoTitle)
}

func main() {
	fmt.Println("===== Observer Pattern Example: YouTube Subscription System =====\n")

	// Create YouTube channels (Subjects)
	techChannel := NewYouTubeChannel("Tech Guru")
	gamingChannel := NewYouTubeChannel("Pro Gamer")

	// Create subscribers (Observers)
	alice := NewSubscriber("Alice")
	bob := NewEmailSubscriber("Bob", "bob@example.com")
	charlie := NewMobileSubscriber("Charlie", "iPhone 15")
	diana := NewSubscriber("Diana")

	fmt.Println("--- Subscribing to channels ---")
	techChannel.Attach(alice)
	techChannel.Attach(bob)
	techChannel.Attach(charlie)

	gamingChannel.Attach(alice)  // Alice likes both tech and gaming
	gamingChannel.Attach(diana)

	// Upload videos
	fmt.Println("\n--- Uploading videos ---")
	techChannel.UploadVideo("10 Best Programming Languages in 2026")

	gamingChannel.UploadVideo("Elden Ring 2 - First Look Gameplay")

	// Unsubscribe
	fmt.Println("\n--- Charlie unsubscribes from Tech Guru ---")
	techChannel.Detach(charlie)

	// Upload another video
	techChannel.UploadVideo("AI Revolution: ChatGPT-5 Released!")

	fmt.Println("\n--- Benefits of Observer Pattern ---")
	fmt.Println("✅ One-to-many notification automatically")
	fmt.Println("✅ Loose coupling between channel and subscribers")
	fmt.Println("✅ Subscribers can join/leave dynamically")
	fmt.Println("✅ Different types of observers (Email, Mobile, etc.)")
}
