namespace DataAccess.Repositories;

public interface IUserRepository
{
    ValueTask<User> CreateAsync(User user, CancellationToken token = default);
    
}