using DataAccess;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;

namespace BusinessLogic;

public static class ServiceExtension
{
    public static void ConfigureCrmServices(this IServiceCollection services, IConfiguration configuration)
    {
        services.ConfigureDataAcces(configuration);

        services.AddTransient<IUserService, UserService>();
        services.AddTransient<ITransactionService, TransactionService>();
        services.AddTransient<IWalletService, WalletService>();
    }
}