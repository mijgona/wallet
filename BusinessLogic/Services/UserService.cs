using DataAccess;
using DataAccess.Repositories;

namespace BusinessLogic;

public sealed class UserService : IUserService
{
    private readonly IUserRepository _userRepository;

    public UserService(IUserRepository userRepository)
    {
        _userRepository = userRepository;
    }

    public async ValueTask<User> CreateUserAsync(UserInfo userInfo, CancellationToken token)
    {
        User newUser = new()
        {
            Name = userInfo.FirstName,
            LastName = userInfo.LastName,
            MiddleName = userInfo.MiddleName,
            Age = userInfo.Age,
            PhoneNumber = userInfo.Phone,
            PassportNumber = userInfo.PassportNumber,
            Gender = userInfo.Gender.ToGenderEnum(),
            Password = userInfo.Password
        };

        return await _userRepository.CreateAsync(newUser, token);;
    }
    
    
    public async ValueTask<User> GetUserByUsernameAsync(string userName, CancellationToken token)
    {
        if (userName is not { Length: > 0 })
            throw new ArgumentNullException(nameof(userName));

        return await _userRepository.GetUserByUserNameAsync(userName, token);
    }

    public async ValueTask<User> GetUserByPhoneAndPasswordAsync(LoginRequest request, CancellationToken token)
    {
        if (request.PhoneNumber is not { Length: > 0 })
            throw new ArgumentNullException(nameof(request.PhoneNumber));
        if (request.Password is not { Length: > 0 })
            throw new ArgumentNullException(nameof(request.Password));

        return await _userRepository.GetUserByPhoneAndPassword(request.PhoneNumber, request.Password, token);
    }
}
