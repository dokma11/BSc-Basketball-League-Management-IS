import { HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { TokenStorage } from "./token.service";

@Injectable()
export class JwtInterceptor implements HttpInterceptor {
  
  constructor(
    private tokenStorage: TokenStorage
  ) {}

  intercept(
    request: HttpRequest<any>,
    next: HttpHandler
  ): Observable<HttpEvent<any>> {
    const accessToken = this.tokenStorage.getAccessToken();
    if (accessToken !== undefined && accessToken !== '' && accessToken !== null) {
      const accessTokenRequest = request.clone({
        setHeaders: {
          Authorization: `Bearer ${accessToken}`,
        },
      });
      return next.handle(accessTokenRequest);
    } else {
      return next.handle(request);
    }
  }
}