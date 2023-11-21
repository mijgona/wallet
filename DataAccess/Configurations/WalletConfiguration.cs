using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace DataAccess;

public sealed class WalletConfiguration: IEntityTypeConfiguration<Wallet>
{
    public void Configure(EntityTypeBuilder<Wallet> modelBuilder)
    {
        modelBuilder
            .HasKey(t => t.Id)
            .HasName("wl_id");
    }
}