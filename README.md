# Gmail-Like Email Application Backend (Golang)

This project is a collaborative effort to build a robust and feature-rich backend for an email application similar to Gmail. We focused on implementing core functionalities like user management, email sending and receiving, drafts, inbox management, and attachment handling. 

## Key Features and Implementations:

**1. User Authentication and Management:**

- **Secure Registration:** Users can register with their name, email address, and password. Email verification is enforced for account activation, ensuring only valid users can access the application.
- **Robust Login:**  Login functionality uses bcrypt for secure password hashing, protecting user credentials.
- **Secure Logout:** Logout is implemented with refresh token blacklisting using Redis, enhancing security by preventing the reuse of compromised tokens.
- **Password Recovery:** A user-friendly password recovery system allows users to reset their passwords using email-based verification codes.

**2. Email Handling:**

- **Comprehensive Email Management:** The backend supports a full range of email operations:
    - **Sending Emails:** Users can compose and send emails with support for To, CC, and BCC recipients.
    - **Receiving Emails:** Incoming emails are received and stored securely in user inboxes.
    - **Inbox Management:** Users can organize their inbox by:
        - Marking emails as read, unread, spam, starred, or archived.
        - Moving emails to and from the trash.
        - Permanently deleting emails.
- **Draft Functionality:**
    - Users can create, update, and delete drafts, providing flexibility in composing emails.
    - Drafts can be sent as complete emails at a later time.

**3. Advanced Email Features:**

- **BCC Handling:**  We implemented a specific solution to ensure that BCC recipients are hidden from To and CC recipients when retrieving emails, preserving privacy as in standard email protocols.
- **Attachment Management:** 
    - Users can upload attachments to Minio, a reliable and scalable object storage service.
    - Attachments are securely associated with outbox messages, ensuring they are sent with the corresponding email.
    - The system allows for efficient retrieval and deletion of attachments.

**4. Performance and Security:**

- **Rate Limiting:**  We implemented rate limiting on sending emails using Redis. This prevents abuse by limiting the number of emails a user can send within a specific time window, ensuring service stability and a fair user experience.
- **Error Logging:**  We integrated a custom logger to capture critical errors and provide detailed information for debugging, making it easier to identify and resolve issues.

## Technical Stack:

- **Programming Language:** Golang (chosen for its performance, concurrency, and strong ecosystem for building APIs and backend systems).
- **Framework:** Gin (a lightweight and fast web framework for Golang, ideal for building RESTful APIs).
- **Database:** PostgreSQL (a powerful and reliable relational database for storing user data, emails, and related information).
- **Object Storage:** Minio (an open-source object storage server compatible with Amazon S3, providing a scalable and cost-effective solution for storing attachments).
- **Caching and Rate Limiting:** Redis (an in-memory data store used for caching frequently accessed data and implementing rate limiting).
- **API Documentation:** Swagger (for automatically generating API documentation, making it easy for developers to understand and use the API).

## Usage:

This application is designed to be run using Docker Compose. Follow these steps to get started:

1. **Clone the repository:** `git clone [repository URL]`
2. **Navigate to the project directory:** `cd [project directory]`
3. **Start the application using Docker Compose:** `docker-compose up -d`

This command will:

- Build the necessary Docker images for the API gateway, Redis, Minio, PostgreSQL, and the Gmail service.
- Create and start the containers for each service.
- Connect the containers to the `gmail_network` network, allowing them to communicate with each other.

4. **Access the API:** The API gateway will be accessible at `http://localhost:8000`. You can use tools like Postman or curl to interact with the API endpoints. Refer to the Swagger documentation for details on the available endpoints and their usage.

5. **Stop the application:** `docker-compose down`

**Note:**

- Make sure you have Docker and Docker Compose installed on your system.
- The `docker-compose.yml` file includes the necessary configurations for each service, including port mappings, environment variables, and dependencies.
- You can customize the `docker-compose.yml` file to adjust the configurations as needed.

## Conclusion:

This project demonstrates a solid foundation for a Gmail-like email application backend. It showcases our collaborative efforts in designing, implementing, and testing core functionalities using a modern and efficient technology stack. The project is well-positioned for future enhancements and can serve as a starting point for building a full-fledged email application.