using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace DataAccess;

public sealed class WalletConfiguration: IEntityTypeConfiguration<WalletBalance>
{
    public void Configure(EntityTypeBuilder<WalletBalance> modelBuilder)
    {
        modelBuilder
            .HasKey(t => t.Id)
            .HasName("pk_id");

        modelBuilder
            .Property(p => p.Id)
            .HasColumnType("SERIAL")
            .HasColumnName("id")
            .IsRequired();

        modelBuilder
            .HasOne(p => p.User)
            .WithOne(w => w.Balance)
            .HasForeignKey<WalletBalance>(w => w.UserId);
    }
}