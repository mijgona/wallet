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
}