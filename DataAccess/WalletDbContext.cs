using Microsoft.EntityFrameworkCore;

namespace DataAccess;

public sealed class WalletDbContext : DbContext
{
    public WalletDbContext(DbContextOptions<WalletDbContext> options)
        : base(options)
    {

    }

    public DbSet<Transaction> Transactions { get; set; }
    public DbSet<User> Users { get; set; }
    public DbSet<Wallet?> Wallet { get; set; }

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.ApplyConfiguration(new UserConfiguration());
        modelBuilder.ApplyConfiguration(new TransactionConfiguration());
        modelBuilder.ApplyConfiguration(new WalletConfiguration());
    }
}
