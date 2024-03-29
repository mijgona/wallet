﻿using System.Text;
using Microsoft.Extensions.Logging;

namespace DataAccess.Repositories;

public sealed class EfCoreUserRepository : IUserRepository
{
    private readonly WalletDbContext _db;

    public EfCoreUserRepository(WalletDbContext walletDbContext)
    {
        _db = walletDbContext;
    }
    
    public async ValueTask<User> CreateAsync(User user, CancellationToken token = default)
    {
        var res = await _db.Users.AddAsync(user, token);
        await _db.SaveChangesAsync(token);
        
        return await _db.Users.FindAsync(res.Entity.Id, token) ?? new User();
    }    
    
    public async ValueTask<User> GetUserByUserNameAsync(string userName, CancellationToken token = default)
    {
        return await ValueTask.FromResult(_db.Users
            .FirstOrDefault(t => t.UserName == userName) ?? throw new InvalidOperationException());
    }

    public async ValueTask<User> GetUserByPhoneAndPassword(string phoneNumber, string password, CancellationToken token)
    {
        return await ValueTask.FromResult(_db.Users
            .FirstOrDefault(t => t.PhoneNumber == phoneNumber && t.Password == password) ?? throw new InvalidOperationException());
    }
}