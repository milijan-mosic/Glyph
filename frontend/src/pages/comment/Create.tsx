import axios from "axios";
import { useState } from "react";

type Props = {
  articleId: string;
  onSuccess: () => void;
};

export const CreateCommentForm = ({ articleId, onSuccess }: Props) => {
  const [author, setAuthor] = useState("");
  const [content, setContent] = useState("");
  const [loading, setLoading] = useState(false);

  const submitComment = async () => {
    if (!author || !content) return;

    setLoading(true);
    try {
      await axios.post("/api/1.0/comment/create", {
        article_id: articleId,
        author_name: author,
        content: content,
      });

      setAuthor("");
      setContent("");
      onSuccess();
    } catch (err) {
      console.error("Failed to submit comment", err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="mt-10 p-6 bg-gray-800 rounded-xl">
      <h3 className="text-lg font-semibold mb-4">Leave a comment</h3>

      <input
        type="text"
        placeholder="Your name"
        className="input w-full mb-3"
        value={author}
        onChange={(e) => setAuthor(e.target.value)}
      />

      <textarea
        placeholder="Your comment"
        className="textarea w-full mb-4"
        rows={4}
        value={content}
        onChange={(e) => setContent(e.target.value)}
      />

      <button
        className="btn btn-primary"
        onClick={submitComment}
        disabled={loading}
      >
        {loading ? "Sending..." : "Submit"}
      </button>

      <p className="text-xs text-gray-400 mt-2">
        Comment will be visible after approval.
      </p>
    </div>
  );
};
