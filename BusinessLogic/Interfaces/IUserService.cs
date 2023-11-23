using DataAccess;

namespace BusinessLogic;

public interface IUserService
{
    public ValueTask<User> CreateUserAsync(UserInfo userInfo, CancellationToken token);
    public ValueTask<User> GetUserByUsernameAsync(string userName, CancellationToken token);
}
