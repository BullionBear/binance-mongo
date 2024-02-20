# Docker build chain

## Authorization
### Public Access
```
aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws/b8k8t3h1
```
### Private Access
```
aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 179335631601.dkr.ecr.ap-northeast-1.amazonaws.com
```

## Base container
### Build
```
docker build -f Dockerfiles/Dockerfile.base -t public.ecr.aws/b8k8t3h1/bmgo:base .
```

### Push
```
docker push public.ecr.aws/b8k8t3h1/bmgo:base
```

