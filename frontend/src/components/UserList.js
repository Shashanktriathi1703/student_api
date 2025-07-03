import React, { useState, useEffect } from 'react';
import { getAllUsers, createUser, deleteUser } from '../services/userService';
import './UserList.css'

function UserList() {
    // State management for users list and new user form
    const [users, setUsers] = useState([]); // यहां हम सभी users का डेटा store करते हैं 😊
    const [newUser, setNewUser] = useState({ name: '', email: '' }); // नए user के लिए form data 📝

    // Component mount पर users fetch करना
    useEffect(() => {
        fetchUsers(); // जैसे ही page load होगा, सारे users fetch हो जाएंगे! 🚀
    }, []);

    // API से users data fetch करने का function
    const fetchUsers = async () => {
        try {
            const fetchedUsers = await getAllUsers(); // API call to get all users
            setUsers(fetchedUsers); // मिले हुए users को state में update कर देते हैं 💯
        } catch (error) {
            console.error('Error fetching users:', error); // अगर कोई error आए तो console में दिखाएंगे 😱
        }
    };

    // नया user create करने का handler
    const handleCreateUser = async (e) => {
        e.preventDefault(); // form submission का default behavior prevent करते हैं
        try {
            const createdUser = await createUser(newUser); // API call करके नया user create करते हैं 🚀
            setUsers([...users, createdUser]); // पुराने users में नया user add कर देते हैं ✅
            setNewUser({ name: '', email: '' }); // form fields को reset कर देते हैं ताकि नया user add कर सकें 🔄
        } catch (error) {
            console.error('Error creating user:', error); // अगर कोई error आए तो console में दिखाएंगे 😱
        }
    };

    // User delete करने का handler function
    const handleDeleteUser = async (id) => {
        try {
            await deleteUser(id); // API call करके user को delete करते हैं ❌
            // Filter method से deleted user को remove करके users state को update करते हैं 🗑️
            setUsers(users.filter(user => user.id !== id));
            // Success message को console में log करते हैं ✅
            console.log(`User with ID ${id} successfully deleted`);
        } catch (error) {
            console.error('Error deleting user:', error); // अगर कोई error आए तो console में दिखाएंगे 😱
            // User को error का notification दे सकते हैं (future enhancement) 🔔
        }
    };

    return (
        <div className="user-management-container">
            <h2 className="user-management-title">User Management</h2>
            <form onSubmit={handleCreateUser} className="user-form">
                <div className="form-group">
                    <input
                        type="text"
                        placeholder="Name"
                        value={newUser.name}
                        onChange={(e) => setNewUser({...newUser, name: e.target.value})}
                        required
                        className="form-input"
                    />
                    <input
                        type="email"
                        placeholder="Email"
                        value={newUser.email}
                        onChange={(e) => setNewUser({...newUser, email: e.target.value})}
                        required
                        className="form-input"
                    />
                    <button type="submit" className="add-button">Add User</button>
                </div>
            </form>
            <div className="table-container">
                <table className="user-table">
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>Name</th>
                            <th>Email</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {users.map(user => (
                            <tr key={user.id}>
                                <td>{user.id}</td>
                                <td>{user.name}</td>
                                <td>{user.email}</td>
                                <td>
                                    <button 
                                        onClick={() => handleDeleteUser(user.id)}
                                        className="delete-button"
                                    >
                                        Delete
                                    </button>
                                </td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
        </div>
    );
}

export default UserList;
