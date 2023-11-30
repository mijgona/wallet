namespace DataAccess.Repositories;

public interface IUserRepository
{
    public ValueTask<User> CreateAsync(User user, CancellationToken token);
    public ValueTask<User> GetUserByUserNameAsync(string userName, CancellationToken token);
    public ValueTask<User> GetUserByPhoneAndPassword(string phoneNumber, string password, CancellationToken token);
}