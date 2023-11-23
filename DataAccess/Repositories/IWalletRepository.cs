namespace DataAccess.Repositories;

public interface IWalletRepository
{
    public ValueTask<Wallet?> CreateAsync(Wallet? wallet, CancellationToken token = default);
    public ValueTask<Wallet?> GetWalletByUserId(long userId, CancellationToken token = default);
    
}