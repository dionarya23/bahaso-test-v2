import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import ArticleList from './ArticleList';

function AdminDashboard() {
  const [articles, setArticles] = useState([]);
  const [page, setPage] = useState(1);
  const [totalPages, setTotalPages] = useState(0);

  useEffect(() => {
    fetchArticles(page);
  }, [page]);

  const fetchArticles = (pageNum) => {
    fetch(`${process.env.REACT_APP_BASE_URL}/api/v1/article/admin?page=${pageNum}`, {
      method: "GET",
      headers: {
        "Authorization": `${JSON.parse(localStorage.getItem("user")).accessToken}`
      }
    })
    .then(response => response.json())
    .then(data => {
      setArticles(data.data.articles);
      setTotalPages(Math.ceil(data.data.total / data.data.limit));
    })
    .catch(error => console.error('Error fetching articles:', error));
  };

  const handleDelete = async (id) => {
    try {
      const response = await fetch(`${process.env.REACT_APP_BASE_URL}/api/v1/article/${id}`, {
        method: 'DELETE',
        headers: {
          "Authorization": `${JSON.parse(localStorage.getItem("user")).accessToken}`
        }
      });
      if (response.ok) {
        setArticles(articles.filter(article => article.id !== id));
      } else {
        console.error('Failed to delete article');
      }
    } catch (error) {
      console.error('Error deleting article:', error);
    }
  };

  return (
    <div className="container mt-4">
      <h2>Admin Dashboard</h2>
      <Link to="/create-article" className="btn btn-success mb-3">Add New Article</Link>
      <ArticleList articles={articles} editable={true} onDelete={handleDelete} />
      {totalPages > 1 && (
        <nav className="d-flex justify-content-center">
          <ul className="pagination">
            {[...Array(totalPages)].map((x, i) => (
              <li key={i + 1} className={`page-item ${i + 1 === page ? 'active' : ''}`}>
                <button className="page-link" onClick={() => setPage(i + 1)}>
                  {i + 1}
                </button>
              </li>
            ))}
          </ul>
        </nav>
      )}
    </div>
  );
}

export default AdminDashboard;
