import React, { useState, useEffect } from 'react';
import ArticleList from './ArticleList';

function Home() {
  const [articles, setArticles] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(0);

  useEffect(() => {
    fetchArticles(currentPage);
  }, [currentPage]);

  const fetchArticles = async (page) => {
    setLoading(true);
    setError(null);
    try {
      const response = await fetch(`${process.env.REACT_APP_BASE_URL}/api/v1/article?page=${page}`);
      if (!response.ok) {
        throw new Error('Failed to fetch articles');
      }
      const data = await response.json();
      setArticles(data.data.articles);
      setTotalPages(Math.ceil(data.data.total / data.data.limit));
    } catch (err) {
      setError('Failed to fetch articles: ' + err.message);
    } finally {
      setLoading(false);
    }
  };

  const handlePageChange = (page) => {
    setCurrentPage(page);
  };

  return (
    <div className="container text-center mt-5">
      <div className="jumbotron">
        <h1 className="display-4">Welcome to Our Blog!</h1>
        <p className="lead">Explore insightful articles on various topics. We publish articles regularly to keep you informed and entertained!</p>
        <hr className="my-4" />
        <p>Check out our latest articles below!</p>
      </div>

      {loading ? (
        <p>Loading articles...</p>
      ) : error ? (
        <p>Error: {error}</p>
      ) : (
        <ArticleList articles={articles} />
      )}

      {totalPages > 1 && (
        <nav aria-label="Article pagination">
          <ul className="pagination justify-content-center">
            {[...Array(totalPages)].map((_, index) => (
              <li key={index} className={`page-item ${index + 1 === currentPage ? 'active' : ''}`}>
                <button className="page-link" onClick={() => handlePageChange(index + 1)}>
                  {index + 1}
                </button>
              </li>
            ))}
          </ul>
        </nav>
      )}
    </div>
  );
}

export default Home;
