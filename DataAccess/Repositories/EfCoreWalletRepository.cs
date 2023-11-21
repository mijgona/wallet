namespace DataAccess.Repositories;

public sealed class EfCoreWalletRepository : IWalletRepository
{
    private readonly WalletDbContext _db;

    public EfCoreWalletRepository(WalletDbContext walletDbContext)
    {
        _db = walletDbContext;
    }
    
    public async ValueTask<Wallet?> CreateAsync(Wallet? wallet, CancellationToken token = default)
    {
        var res = await _db.Wallet.AddAsync(wallet, token);
        await _db.SaveChangesAsync(token);

        if (res.Entity != null) return await _db.Wallet.FindAsync(res.Entity.Id, token);
        return new Wallet();
    }

    public ValueTask<Wallet?> GetWalletByUserId(long userId, CancellationToken token = default)
    {
        var res = _db.Wallet
            .FirstOrDefault(t => t.UserId == userId);
        return ValueTask.FromResult(res);
    }
}