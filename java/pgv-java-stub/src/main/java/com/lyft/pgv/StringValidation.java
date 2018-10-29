package com.lyft.pgv;

import java.util.Arrays;

import org.apache.commons.validator.routines.DomainValidator;
import org.apache.commons.validator.routines.InetAddressValidator;
import org.apache.commons.validator.routines.EmailValidator;
import org.apache.commons.validator.routines.UrlValidator;
import com.google.re2j.*;

public final class StringValidation {
    private StringValidation() { }

    public static void constant(String field, String value, String expected) throws ValidationException {
        if (!value.equals(expected)) {
            throw new ValidationException(field, "value must equal " + expected);
        }
    }

    public static void in(String field, String value, String[] set) throws ValidationException {
        for (String str : set) {
            if (value.equals(str)) {
                return;
            }
        }

        throw new ValidationException(field, "value must be in " + Arrays.toString(set));
    }

    public static void notIn(String field, String value, String[] set) throws ValidationException {
        for (String str : set) {
            if (value.equals(str)) {
                throw new ValidationException(field, "value must not be in " + Arrays.toString(set));
            }
        }
    }

    public static void length(String field, String value, int expected) throws ValidationException {
        if (value.length() != expected) {
                throw new ValidationException(field, "length must be " + expected);
        }
    }

    public static void minLength(String field, String value, int expected) throws ValidationException {
        if (value.length() < expected) {
                throw new ValidationException(field, "length must be at least " + expected);
        }
    }

    public static void maxLength(String field, String value, int expected) throws ValidationException {
        if (value.length() > expected) {
                throw new ValidationException(field, "length must be at maximum " + expected);
        }
    }

    public static void lenBytes(String field, String value, int expected) throws ValidationException {
        if (value.getBytes().length != expected) {
                throw new ValidationException(field, "bytes length must be " + expected);
        }
    }

    public static void minBytes(String field, String value, int expected) throws ValidationException {
        if (value.getBytes().length < expected) {
                throw new ValidationException(field, "bytes length must be at least " + expected);
        }
    }

    public static void maxBytes(String field, String value, int expected) throws ValidationException {
        if (value.getBytes().length > expected) {
                throw new ValidationException(field, "bytes length must be at maximum " + expected);
        }
    }

    public static void pattern(String field, String value, String pattern) throws ValidationException {
      Pattern p = Pattern.compile(pattern);
      if (!p.matches(value)) {
               throw new ValidationException(field, "must match pattern " + pattern);
      }
    }

    public static void prefix(String field, String value, String prefix) throws ValidationException {
      if (!value.startsWith(prefix)) {
               throw new ValidationException(field, "should start with " + prefix);
      }
    }

    public static void contains(String field, String value, String contains) throws ValidationException {
      if (!value.contains(contains)) {
               throw new ValidationException(field, "should contain " + contains);
      }
    }

    public static void suffix(String field, String value, String suffix) throws ValidationException {
      if (!value.endsWith(suffix)) {
               throw new ValidationException(field, "should end with " + suffix);
      }
    }

    public static void email(String field, String value) throws ValidationException {
      EmailValidator emailValidator = EmailValidator.getInstance();
      if (!emailValidator.isValid(value)) {
               throw new ValidationException(field, "should be valid email " + value);
      }
    }

    public static void hostName(String field, String value) throws ValidationException {
      DomainValidator domainValidator = DomainValidator.getInstance();
      if (!domainValidator.isValid(value)) {
               throw new ValidationException(field, "should be valid host " + value);
      }
    }

    public static void ip(String field, String value) throws ValidationException {
      InetAddressValidator ipValidator = InetAddressValidator.getInstance();
      if (!ipValidator.isValid(value)) {
               throw new ValidationException(field, "should be valid ip address " + value);
      }
    }

    public static void ipv4(String field, String value) throws ValidationException {
      InetAddressValidator ipValidator = InetAddressValidator.getInstance();
      if (!ipValidator.isValidInet4Address(value)) {
               throw new ValidationException(field, "should be valid ipv4 address " + value);
      }
    }

    public static void ipv6(String field, String value) throws ValidationException {
      InetAddressValidator ipValidator = InetAddressValidator.getInstance();
      if (!ipValidator.isValidInet6Address(value)) {
               throw new ValidationException(field, "should be valid ipv6 address " + value);
      }
    }

    public static void uri(String field, String value) throws ValidationException {
      UrlValidator urlValidator = UrlValidator.getInstance();
      if (!urlValidator.isValid(value)) {
               throw new ValidationException(field, "should be valid uri " + value);
      }
    }

    public static void uriRef(String field, String value) throws ValidationException {
      UrlValidator urlValidator = UrlValidator.getInstance();
      if (!urlValidator.isValid(value)) {
               throw new ValidationException(field, "should be valid uri ref " + value);
      }
    }
}