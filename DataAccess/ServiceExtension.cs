using DataAccess;
using DataAccess.Repositories;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Design;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;

namespace DataAccess;

public static class ServiceExtension
{
    private static string DefaultConnectionKeyName => "DefaultConnection";
    public static void ConfigureDataAcces(this IServiceCollection services, IConfiguration configuration)
    {
        services.AddDbContext<WalletDbContext>(
            opt => opt.UseNpgsql(configuration.GetConnectionString(DefaultConnectionKeyName)));
        services.AddScoped<ITransactionRepository, EfCoreTransactionRepository>();
        services.AddScoped<IUserRepository, EfCoreUserRepository>();
        services.AddScoped<IWalletRepository, EfCoreWalletRepository>();

    }
}

public class WalletDbContextFactory : IDesignTimeDbContextFactory<WalletDbContext>
{
    public WalletDbContext CreateDbContext(string[] args)
    {
        var optionsBuilder = new DbContextOptionsBuilder<WalletDbContext>();
        optionsBuilder.UseNpgsql("Server=localhost;Port=5432;User ID=postgres;Password=changeme;Database=wallet;");

        return new WalletDbContext(optionsBuilder.Options);
    }
}