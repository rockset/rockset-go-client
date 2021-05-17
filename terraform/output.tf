output "rockset-role-arn" {
  value       = aws_iam_role.rockset.arn
  description = "AWS of Rockset integration role"
}
