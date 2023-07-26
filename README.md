# AImate - Advanced Knowledge Extraction and Question-Answering Tool

## Introduction

AImate is an advanced knowledge extraction and question-answering tool that leverages cutting-edge technologies like OpenAI and Pinecone Vector Database. This powerful tool enables users to upload their custom knowledgebase files and obtain accurate and contextually relevant answers to their queries. The primary focus is on processing human-readable content such as books, letters, and documents, making AImate an indispensable tool for extracting knowledge and insights from vast repositories of textual data.

## Key Features

- **Custom Knowledge Base:** Users can effortlessly upload various document types, including PDFs, plain text, .docx, and .epub files, to create their custom knowledge base.

- **Intelligent Question-Answering:** AImate utilizes state-of-the-art AI from OpenAI to provide precise and informative answers to user questions based on the content within the uploaded documents.

- **Contextual Snippets:** The system not only presents answers but also provides context snippets from the source documents, allowing users to understand the reasoning behind the answers.

- **Scalable Database:** AImate harnesses the Pinecone Vector Database to efficiently store and index embeddings generated from the uploaded content. This ensures fast and accurate retrieval of information during question-answering.

## Tech Stack

AImate is built using a robust tech stack, combining the power of various technologies to deliver seamless knowledge extraction and question-answering capabilities:

- **Frontend:** The user-friendly and interactive frontend is developed using React.js, a popular JavaScript library for building modern web applications.

- **Backend:** The backend server is built using Golang (Go), a high-performance language known for its concurrency support and efficiency.

- **Embedding Generation:** AImate employs OpenAI's state-of-the-art Natural Language Processing models to generate meaningful embeddings for text data, forming the basis of accurate question-answering.

- **Storage and Retrieval:** The Pinecone Vector Database is integrated into the system to efficiently store the embeddings, enabling fast and scalable retrieval of information during question-answering.

- **Text Extraction:** For document parsing and text extraction, AImate uses Poppler, a widely used PDF and document utility.

## Future Development: Dockerization and Kubernetes Deployment

In future development, AImate is planned to be dockerized to simplify deployment across different environments. Docker containers will encapsulate the application and its dependencies, ensuring consistency and reproducibility during deployment. Additionally, AImate aims to leverage Kubernetes for container orchestration, enabling seamless scaling and management of the application in a distributed environment. Kubernetes will empower AImate to handle high loads and provide a robust and fault-tolerant user experience.

## Setup

To set up AImate on your local machine, follow these steps:

1. Ensure you have Node v19 and Go v1.18.9 (darwin/arm64) installed.

2. Install the necessary dependencies:
   ```
   sudo apt-get install -y poppler-utils # On Ubuntu
   brew install poppler # On Mac
   ```

3. Obtain API keys and endpoints:
   - Create a new file `secret/openai_api_key` and paste your OpenAI API key into it.
   - Create a new file `secret/pinecone_api_key` and paste your Pinecone API key into it.
   - Create a new file `secret/pinecone_api_endpoint` and paste your Pinecone API endpoint URL into it.

4. Install JavaScript package dependencies:
   ```
   npm install
   ```

5. Start the Golang web server:
   ```
   npm start
   ```

6. Compile the JS code and create a `bundle.js` file:
   ```
   npm run dev
   ```

7. Visit the local version of the site at [http://localhost:8100](http://localhost:8100)

## Under the Hood

AImate's backend server handles incoming uploads and question-answering using POST APIs:

- `/upload`: Handles file uploads, processing them into embeddings stored in the Pinecone Vector Database.
- `/api/question`: Handles incoming questions, utilizing OpenAI embeddings to retrieve the most relevant answers from the database.

## Generative Question-Answering with Long-Term Memory

For more in-depth insights into the generative question-answering technique employed by AImate, we recommend reading the post from the Pinecone blog: [Generative Question-Answering with Long-Term Memory](https://www.pinecone.io/learn/openai-gen-qa/).

We hope you enjoy using AImate for knowledge extraction and question-answering. For any inquiries or feedback, feel free to reach out!
