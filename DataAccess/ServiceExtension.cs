using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;

namespace DataAccess;

public static class ServiceExtension
{
    private static string DefaultConnectionKeyName => "DefaultConnection";
    public static void ConfigureDataAccess(this IServiceCollection services, IConfiguration configuration)
    {
        services.AddDbContext<WalletDbContext>(
            opt => opt.UseNpgsql(configuration.GetConnectionString(DefaultConnectionKeyName)));

        services.AddScoped<IOrderRepository, EfCoreOrderRepository>();
    }
}