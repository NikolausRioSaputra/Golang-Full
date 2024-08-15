// implementasi metode yang digunakan untuk mendapatkan semua user
// func (repo *UserRepository) GetAllUser(ctx context.Context) ([]domain.User, error) {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()
// 	var users []domain.User
// 	db := db.GetDB()
// 	rows, err := db.QueryContext(ctx, `SELECT id, name, address, balance FROM user`)
// 	if err != nil {
// 		return users, err
// 	}
// 	for rows.Next() {
// 		var user domain.User
// 		err := rows.Scan(&user.ID, &user.Name, &user.Address, &user.Balance)
// 		if err != nil {
// 			return users, err
// 		}
// 		users = append(users, user)
// 	}

// 	return users, nil
// }