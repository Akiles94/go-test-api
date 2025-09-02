# scripts/test.ps1
Write-Host "🧪 Testing microservices locally..." -ForegroundColor Green

Write-Host "=== CONNECTIVITY TESTS ===" -ForegroundColor Cyan
Write-Host "1️⃣ Testing domain resolution:" -ForegroundColor Yellow
nslookup api.local.dev

Write-Host "`n2️⃣ Basic connectivity:" -ForegroundColor Yellow
try {
    $response = Invoke-WebRequest -Uri "http://api.local.dev/health" -UseBasicParsing -TimeoutSec 10
    Write-Host "✅ Health endpoint responding: $($response.StatusCode)" -ForegroundColor Green
} catch {
    Write-Host "❌ Health endpoint not responding: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n=== USER SERVICE TESTS ===" -ForegroundColor Cyan
Write-Host "3️⃣ Register a new user:" -ForegroundColor Yellow

$registerBody = @{
    name = "Test"
    lastName = "User"
    email = "test@local.dev"
    password = "Password123!"
    role = 1
} | ConvertTo-Json

try {
    $registerResponse = Invoke-RestMethod -Uri "http://api.local.dev/api/v1/auth/register" -Method POST -Body $registerBody -ContentType "application/json"
    Write-Host "Register response: $($registerResponse | ConvertTo-Json)" -ForegroundColor Green
} catch {
    Write-Host "Register failed: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n4️⃣ Login with the user:" -ForegroundColor Yellow
$loginBody = @{
    email = "test@local.dev"
    password = "Password123!"
} | ConvertTo-Json

try {
    $loginResponse = Invoke-RestMethod -Uri "http://api.local.dev/api/v1/auth/login" -Method POST -Body $loginBody -ContentType "application/json"
    $token = $loginResponse.token
    Write-Host "Login successful! Token: $($token.Substring(0, 50))..." -ForegroundColor Green
    
    Write-Host "`n=== PRODUCT SERVICE TESTS ===" -ForegroundColor Cyan
    Write-Host "5️⃣ Try accessing categories without token (should fail):" -ForegroundColor Yellow
    
    try {
        Invoke-RestMethod -Uri "http://api.local.dev/api/v1/categories" -Method GET
    } catch {
        Write-Host "✅ Correctly rejected request without token" -ForegroundColor Green
    }
    
    if ($token) {
        Write-Host "`n6️⃣ Access categories with token:" -ForegroundColor Yellow
        $headers = @{ Authorization = "Bearer $token" }
        
        try {
            $categories = Invoke-RestMethod -Uri "http://api.local.dev/api/v1/categories" -Method GET -Headers $headers
            Write-Host "Categories: $($categories | ConvertTo-Json)" -ForegroundColor Green
        } catch {
            Write-Host "Failed to get categories: $($_.Exception.Message)" -ForegroundColor Red
        }
    }
    
} catch {
    Write-Host "Login failed: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n✅ Testing complete!" -ForegroundColor Green