namespace DataAccess.Repositories;

public interface IUserRepository
{
    public ValueTask<User> CreateAsync(User user, CancellationToken token);
    public ValueTask<User> GetUserByUserName(string userName, CancellationToken token);

}