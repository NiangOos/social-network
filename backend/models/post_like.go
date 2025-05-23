package models

import "database/sql"

// PostLike structure represents the "post_likes" table
type PostLike struct {
	PostLikeID int `json:"post_like_id"`
	AuthorID   int `json:"author_id"`
	PostID     int `json:"post_id"`
	Rate       int `json:"rate"`
}

type PostLikeRepository struct {
	db *sql.DB
}

func NewPostLikeRepository(db *sql.DB) *PostLikeRepository {
	return &PostLikeRepository{
		db: db,
	}
}

// CreatePostLike adds a new post like to the database
func (plr *PostLikeRepository) CreatePostLike(postLike *PostLike) error {
	query := `
		INSERT INTO post_likes (author_id, post_id, rate)
		VALUES (?, ?, ?)
	`
	_, err := plr.db.Exec(query, postLike.AuthorID, postLike.PostID, postLike.Rate)
	if err != nil {
		return err
	}

	return nil
}

// GetPostLike retrieves a post like from the database by post_like_id
func (plr *PostLikeRepository) GetPostLike(postLikeID int) (*PostLike, error) {
	query := "SELECT * FROM post_likes WHERE post_like_id = ?"
	var postLike PostLike
	err := plr.db.QueryRow(query, postLikeID).Scan(&postLike.PostLikeID, &postLike.AuthorID, &postLike.PostID, &postLike.Rate)
	if err != nil {
		return nil, err
	}
	return &postLike, nil
}
// GetPostLike retrieves a post like from the database by post_like_id
// func  (plr *PostLikeRepository) IsLiked(postID, user_id int) (*PostLike, error) {
// 	query := "SELECT * FROM post_likes WHERE PostID = ? AND AuthorID = ?"
// 	var postLike PostLike
// 	err := plr.db.QueryRow(query, postID, user_id).Scan(&postLike.PostLikeID, &postLike.AuthorID, &postLike.PostID, &postLike.Rate)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &postLike, nil
// }

// GetNumberOfLikes retrieves the number of likes for a given post ID
func (plr *PostLikeRepository) GetNumberOfLikes(postID int) (int, error) {
	query := "SELECT COUNT(*) FROM post_likes WHERE post_id = ?"
	var numLikes int
	err := plr.db.QueryRow(query, postID).Scan(&numLikes)
	if err != nil {
		return 0, err
	}
	return numLikes, nil
}

// IsPostLikedByCurrentUser checks if the current user has liked a given post
func (plr *PostLikeRepository) IsPostLikedByCurrentUser(postID, userID int) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM post_likes WHERE post_id = ? AND author_id = ?)"
	var exists bool
	err := plr.db.QueryRow(query, postID, userID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// UpdatePostLike updates an existing post like in the database
func (plr *PostLikeRepository) UpdatePostLike(postLike *PostLike) error {
	query := `
		UPDATE post_likes
		SET author_id = ?, post_id = ?, rate = ?
		WHERE post_like_id = ?
	`
	_, err := plr.db.Exec(query, postLike.AuthorID, postLike.PostID, postLike.Rate, postLike.PostLikeID)
	if err != nil {
		return err
	}
	return nil
}

// GetUserLikedPosts retrieves posts liked by a specific user from the database
func (pr *PostRepository) GetUserLikedPosts(userID int) ([]*Post, error) {
	rows, err := pr.db.Query("SELECT p.post_id, p.title, p.category, p.content, p.created_at, p.author_id, p.image_url, p.visibility FROM post p INNER JOIN post_like pl ON p.post_id = pl.post_id WHERE pl.author_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.PostID, &post.Title, &post.Category, &post.Content, &post.CreatedAt, &post.AuthorID, &post.Visibility)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	return posts, nil
}

// DeletePostLike removes a post like from the database by post_like_id
func (plr *PostLikeRepository) DeletePostLike(postID, AuthorID int) error {
	query := "DELETE FROM post_likes WHERE post_id = ? AND author_id= ?"
	_, err := plr.db.Exec(query, postID, AuthorID)
	if err != nil {
		return err
	}
	return nil
}

// CountLikesForPost retrieves the number of likes for a given post_id
func (plr *PostLikeRepository) CountLikesForPost(postID int) (int, error) {
	query := `
		SELECT COUNT(*) FROM post_likes
		WHERE post_id = ?
	`

	var count int
	err := plr.db.QueryRow(query, postID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}